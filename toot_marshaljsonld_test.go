package tootns_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-tootns"
)

func TestToot_MarshalJSONLD(t *testing.T) {

	const context string =
	`"@context":{`+
		`"toot":"http://joinmastodon.org/ns#"`+
		`,`+
		`"attributionDomains":"toot:attributionDomains"`+
		`,`+
		`"blurhash":"toot:blurhash"`+
		`,`+
		`"discoverable":"toot:discoverable"`+
//		`,`+
//		`"Emoji":"toot:Emoji"`+
		`,`+
		`"featured":"toot:featured"`+
		`,`+
		`"featuredTags":"toot:featuredTags"`+
		`,`+
		`"focalPoint":"toot:focalPoint"`+
//		`,`+
//		`"IdentityProof":"toot:IdentityProof"`+
		`,`+
		`"indexable":"toot:indexable"`+
		`,`+
		`"memorial":"toot:memorial"`+
		`,`+
		`"suspended":"toot:suspended"`+
		`,`+
		`"votersCount":"toot:votersCount"`+
	`}`

	tests := []struct{
		Value tootns.Toot
		Expected []byte
	}{
		{
			Value: tootns.Toot{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: tootns.Toot{
				AttributionDomains: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"attributionDomains":"apple banana cherry"}`),
		},
		{
			Value: tootns.Toot{
				BlurHash: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"blurhash":"apple banana cherry"}`),
		},
		{
			Value: tootns.Toot{
				Discoverable: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"discoverable":true}`),
		},
		{
			Value: tootns.Toot{
				Discoverable: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"discoverable":false}`),
		},
//		{
//			Value: tootns.Toot{
//				Emoji: opt.Something("apple banana cherry"),
//			},
//			Expected: []byte(`{`+context+`,"Emoji":"apple banana cherry"}`),
//		},
//		{
//			Value: tootns.Toot{
//				Emoji: opt.Something(""),
//			},
//			Expected: []byte(`{`+context+`,"Emoji":""}`),
//		},
		{
			Value: tootns.Toot{
				Featured: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"featured":true}`),
		},
		{
			Value: tootns.Toot{
				Featured: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"featured":false}`),
		},
		{
			Value: tootns.Toot{
				FeaturedTags: []string{"apple","banana","cherry"},
			},
			Expected: []byte(`{`+context+`,"featuredTags":["apple","banana","cherry"]}`),
		},
		{
			Value: tootns.Toot{
				FocalPoint: []any{12,456},
			},
			Expected: []byte(`{`+context+`,"focalPoint":[12,456]}`),
		},
//		{
//			Value: tootns.Toot{
//				IdentityProof: opt.Something("apple banana cherry"),
//			},
//			Expected: []byte(`{`+context+`,"IdentityProof":"apple banana cherry"}`),
//		},
		{
			Value: tootns.Toot{
				Indexable: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"indexable":true}`),
		},
		{
			Value: tootns.Toot{
				Indexable: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"indexable":false}`),
		},
		{
			Value: tootns.Toot{
				Memorial: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"memorial":true}`),
		},
		{
			Value: tootns.Toot{
				Memorial: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"memorial":false}`),
		},
		{
			Value: tootns.Toot{
				VotersCount: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"votersCount":"apple banana cherry"}`),
		},
		{
			Value: tootns.Toot{
				Suspended: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"suspended":true}`),
		},
		{
			Value: tootns.Toot{
				Suspended: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"suspended":false}`),
		},



		{
			Value: tootns.Toot{
				Discoverable: opt.Something(false),
//				Emoji: opt.Something("apple banana cherry"),
				Indexable: opt.Something(true),
				Memorial: opt.Something(false),
				Suspended: opt.Something(false),
			},
//			Expected: []byte(`{`+context+`,"discoverable":false,"Emoji":"apple banana cherry","indexable":true,"memorial":false,"suspended":false}`),
			Expected: []byte(`{`+context+`,"discoverable":false,"indexable":true,"memorial":false,"suspended":false}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual marshaled-jsonld it not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d) %q", len(expected), expected)
				t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
