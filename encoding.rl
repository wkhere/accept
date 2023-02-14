package accept

import (
	"strconv"
	"fmt"
)

type encodingSpec struct {
	token string
	weight float64
}

// See https://www.rfc-editor.org/rfc/rfc7231#section-5.3.4 ,
//     https://www.rfc-editor.org/rfc/rfc7231#section-5.3.1 ,
//     https://www.rfc-editor.org/rfc/rfc7231#appendix-D
// and https://www.rfc-editor.org/rfc/rfc7230#section-3.2.6 .
func parseEncoding(data string) (e encodingSpec, _ error) {
	%% machine encoding;
	%% write data;

	cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := 0

	text := func() string { return data[pb:p] }

	e = encodingSpec{weight: 1}

    %%{
    	action mark { pb = p }

		action setToken  { e.token = text() }
		action setWeight { e.weight, _ = strconv.ParseFloat(text(), 64) }

		ows = (' ' | '\t')*;

		token = (alnum | '!' | '#' | '$' | '%' | '&' | '\'' | '*' |
		                 '+' | '-' | '.' | '^' | '_' | '`'  | '|' | '~')+
				>mark %setToken;

		qvalue  = ( '0' ('.' digit+)? ) | ( '1' ('.' '0'+)? );
		# ^^ fractional part should be only 3-digit as in rfc7231

		weight  = ows ';' ows 'q=' qvalue >mark %setWeight;
		coding  = (token | '*' | 'identity') weight?;

		main := ows coding? ows;

		write init;
		write exec;
    }%%

    if cs < encoding_first_final {
		return e, fmt.Errorf("cs=%d data=%q p=%d", cs, data, p)
    }
    return e, nil
}
