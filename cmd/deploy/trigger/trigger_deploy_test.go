package trigger

import (
	"testing"
)

func TestDecodePk(t *testing.T) {
	// fake private key
	pem := `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCCzfTN2Wccm/XHg62oMnDgKZhUSQExxhKFTt2ofCmXYfBo/yd7
ecvo2qMZB12OTVBPugksODdoSpjGRQUVzaltG6y42tK+dWW/7h6fu4jlDAYuLxta
i3tDjrSYavrjQkzd5BGGb+aPnj0nzV90Ma5tRXl2nfqTfpDILTYMb55slwIDAQAB
AoGAZJpND0mHIZzeEr43AOcSo4W6UBD6JzMFKZx3pM8iGflMsySTVEdfyk7mJCjj
DsBt5XPa/DScgMzm5Y9OEu/jvxhrPvks5ZJcPDMdUgl53CdJzMaFXDEKQVzi9caW
jdOL+nGpOU6NH7xPxSelRBqmx/UGRIWoRCl8Wzo6P2UhDTkCQQDO+6lq8A2SLT5x
NVv3gfrQlYrjtbDOK37GfuL557cdOYjdvcfRWatgtCEj2G17gT6mB79OKNa4MzgB
icNrb2DzAkEAocf5YJ5+SRvWfwLd0ZvAvLUYc7pIja9SARDmHV5WZTZdvJ6KiWjy
bv7nbiDGV7Hs2450XiHxlgwyOVPf032OzQJBAJ1ap1ab/sU1TjZEiZElqKcwOPGa
LDrgyDVhO28fYI+AMPPalnEqiWYwKs2UeM8e16DjXOWvmGVa0uYgdIdVmk8CQDvq
pCF0kbDh7n77wMtws+Ev8O5rf5N56xUZ/R/DYJv7lpvU29ooVCFnpq7S1KKF8wMd
r1ttltvLiI5S0gKx7cECQGc6a6mhdy3Lewmhufs+EfczJ2Qvf3ql5NMr6+MniEmk
n6R5zzzlRy/nxzAOimq6OXoqqQQ++4gELFw/+vMW7rU=
-----END RSA PRIVATE KEY-----`
	_, err := decodePk(pem)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
