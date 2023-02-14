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
func parseEncodings(data string) (ee []encodingSpec, _ error) {
	%% machine encodings;
	%% write data;

	cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := 0

	text := func() string { return data[pb:p] }

	var e encodingSpec
	var pushed bool

    %%{
    	action mark { pb = p }

		action new  {
			e = encodingSpec{weight: 1}
			pushed = false
		}
		action push {
			ee = append(ee, e)
			pushed = true
		}
		action pop {
			if pushed {
				ee = ee[:len(ee)-1]
			}
			fhold; fgoto next;
		}
		action skip { fhold; fgoto next; }

		action setToken  { e.token = text() }
		action setWeight { e.weight, _ = strconv.ParseFloat(text(), 64) }

		ows = (' ' | '\t')*;

		token = (alnum | '!' | '#' | '$' | '%' | '&' | '\'' | '*' |
		                 '+' | '-' | '.' | '^' | '_' | '`'  | '|' | '~')+
				>mark %setToken;

		qvalue  = ( '0' ('.' digit+)? ) | ( '1' ('.' '0'+)? );
		# ^^ fractional part should be only 3-digit as in rfc7231

		weight  = ows ';' ows 'q=' qvalue >mark @err(skip) %setWeight;
		coding  = (token | '*' | 'identity') >new weight? %push @err(pop);

		entry  = ows coding? ows;
		entryN = entry;

		main := entry? (',' entryN)*;

		next := [^,]* ',' @{ fgoto entryN; };

		write init;
		write exec;
    }%%

    if cs < encodings_first_final {
		return ee, fmt.Errorf("cs=%d data=%q p=%d", cs, data, p)
    }
    return ee, nil
}
