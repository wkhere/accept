
//line encoding.rl:1
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
	
//line encoding.rl:19
	
//line encoding.go:24
var _encoding_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 2, 
	0, 3, 0, 3, 3, 3, 
}

const encoding_start int = 6
const encoding_first_final int = 6
const encoding_error int = 0

const encoding_en_main int = 6


//line encoding.rl:20

	cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := 0

	text := func() string { return data[pb:p] }

	e = encodingSpec{weight: 1}

    
//line encoding.go:48
	{
	cs = encoding_start
	}

//line encoding.go:53
	{
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 6:
		switch data[p] {
		case 9:
			goto tr8;
		case 32:
			goto tr8;
		case 33:
			goto tr9;
		case 124:
			goto tr9;
		case 126:
			goto tr9;
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 39:
				if 42 <= data[p] && data[p] <= 43 {
					goto tr9;
				}
			case data[p] >= 35:
				goto tr9;
			}
		case data[p] > 46:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr9;
				}
			case data[p] > 90:
				if 94 <= data[p] && data[p] <= 122 {
					goto tr9;
				}
			default:
				goto tr9;
			}
		default:
			goto tr9;
		}
		goto tr1;
	case 0:
		goto _out
	case 7:
		switch data[p] {
		case 9:
			goto tr10;
		case 32:
			goto tr10;
		case 33:
			goto tr11;
		case 59:
			goto tr12;
		case 124:
			goto tr11;
		case 126:
			goto tr11;
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 39:
				if 42 <= data[p] && data[p] <= 43 {
					goto tr11;
				}
			case data[p] >= 35:
				goto tr11;
			}
		case data[p] > 46:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr11;
				}
			case data[p] > 90:
				if 94 <= data[p] && data[p] <= 122 {
					goto tr11;
				}
			default:
				goto tr11;
			}
		default:
			goto tr11;
		}
		goto tr1;
	case 8:
		switch data[p] {
		case 9:
			goto tr13;
		case 32:
			goto tr13;
		case 59:
			goto tr0;
		}
		goto tr1;
	case 1:
		switch data[p] {
		case 9:
			goto tr0;
		case 32:
			goto tr0;
		case 113:
			goto tr2;
		}
		goto tr1;
	case 2:
		if data[p] == 61 {
			goto tr3;
		}
		goto tr1;
	case 3:
		switch data[p] {
		case 48:
			goto tr4;
		case 49:
			goto tr5;
		}
		goto tr1;
	case 9:
		switch data[p] {
		case 9:
			goto tr14;
		case 32:
			goto tr14;
		case 46:
			goto tr15;
		}
		goto tr1;
	case 10:
		switch data[p] {
		case 9:
			goto tr16;
		case 32:
			goto tr16;
		}
		goto tr1;
	case 4:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6;
		}
		goto tr1;
	case 11:
		switch data[p] {
		case 9:
			goto tr14;
		case 32:
			goto tr14;
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6;
		}
		goto tr1;
	case 12:
		switch data[p] {
		case 9:
			goto tr14;
		case 32:
			goto tr14;
		case 46:
			goto tr17;
		}
		goto tr1;
	case 5:
		if data[p] == 48 {
			goto tr7;
		}
		goto tr1;
	case 13:
		switch data[p] {
		case 9:
			goto tr14;
		case 32:
			goto tr14;
		case 48:
			goto tr7;
		}
		goto tr1;
	}

	tr1: cs = 0; goto _again
	tr0: cs = 1; goto _again
	tr12: cs = 1; goto f1
	tr2: cs = 2; goto _again
	tr3: cs = 3; goto _again
	tr15: cs = 4; goto _again
	tr17: cs = 5; goto _again
	tr8: cs = 6; goto _again
	tr11: cs = 7; goto _again
	tr9: cs = 7; goto f0
	tr13: cs = 8; goto _again
	tr10: cs = 8; goto f1
	tr4: cs = 9; goto f0
	tr16: cs = 10; goto _again
	tr14: cs = 10; goto f2
	tr6: cs = 11; goto _again
	tr5: cs = 12; goto f0
	tr7: cs = 13; goto _again

f0:
//line encoding.rl:30
 pb = p 
	goto _again
f1:
//line encoding.rl:32
 e.token = text() 
	goto _again
f2:
//line encoding.rl:33
 e.weight, _ = strconv.ParseFloat(text(), 64) 
	goto _again

_again:
	if cs == 0 {
		goto _out
	}
	if p++; p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		switch _encoding_eof_actions[cs] {
		case 2:
//line encoding.rl:32
 e.token = text() 
		case 3:
//line encoding.rl:33
 e.weight, _ = strconv.ParseFloat(text(), 64) 
//line encoding.go:289
		}
	}

	_out: {}
	}

//line encoding.rl:51


    if cs < encoding_first_final {
		return e, fmt.Errorf("cs=%d data=%q p=%d", cs, data, p)
    }
    return e, nil
}
