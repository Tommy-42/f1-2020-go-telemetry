package elastic

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/pkg/errors"
)

type Elastic struct {
	client *elasticsearch.Client
	index  string
}

func NewES(conf Config) (*Elastic, error) {

	cfg := elasticsearch.Config{
		Addresses: conf.Addresses,
		Username:  conf.Username,
		Password:  conf.Password,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not start elasticsearch client")
	}

	es := &Elastic{
		client: client,
		index:  conf.Index,
	}

	res, err := client.Indices.Exists([]string{conf.Index})
	if err != nil {
		return nil, errors.Wrapf(err, "could not check elasticsearch index: %s", conf.Index)
	}
	if res.IsError() {
		if res.StatusCode == http.StatusNotFound {
			ret, err := client.Indices.Create(conf.Index)
			if err != nil {
				return nil, errors.Wrapf(err, "could not create elasticsearch index: %s", conf.Index)
			}
			if ret.IsError() {
				return nil, errors.New(fmt.Sprintf("could not create elasticsearch index %s: [%d]%s", conf.Index, ret.StatusCode, ret.Status()))
			}
			return es, nil
		}
		return nil, errors.New(fmt.Sprintf("could not request exists elasticsearch index %s: [%d]%s", conf.Index, res.StatusCode, res.Status()))
	}

	return es, nil
}

func (e *Elastic) Store(ctx context.Context, body *bytes.Reader) error {
	if e == nil {
		return errors.New("elastic is not initialized")
	}

	req := esapi.IndexRequest{
		Index:   e.index,
		Body:    body,
		Timeout: time.Second * 10,
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return errors.Wrap(err, "could not request elasticsearch")
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Wrap(errors.New(res.String()), "could not store packet to elasticsearch")
	}

	return nil
}
