
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
func parseEncodings(data string) (ee []encodingSpec, _ error) {
	
//line encoding.rl:19
	
//line encoding.go:24
var _encodings_eof_actions []byte = []byte{
	0, 1, 1, 2, 2, 2, 0, 0, 
	6, 0, 8, 0, 8, 8, 8, 0, 
}

const encodings_start int = 7
const encodings_first_final int = 7
const encodings_error int = 0

const encodings_en_main int = 7
const encodings_en_main_entryN int = 7
const encodings_en_next int = 6


//line encoding.rl:20

	cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := 0

	text := func() string { return data[pb:p] }

	var e encodingSpec
	var pushed bool

    
//line encoding.go:51
	{
	cs = encodings_start
	}

//line encoding.go:56
	{
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 7:
		switch data[p] {
		case 9:
			goto tr11;
		case 32:
			goto tr11;
		case 33:
			goto tr13;
		case 44:
			goto tr11;
		case 124:
			goto tr13;
		case 126:
			goto tr13;
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 39:
				if 42 <= data[p] && data[p] <= 46 {
					goto tr13;
				}
			case data[p] >= 35:
				goto tr13;
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 94 <= data[p] && data[p] <= 122 {
					goto tr13;
				}
			case data[p] >= 65:
				goto tr13;
			}
		default:
			goto tr13;
		}
		goto tr12;
	case 0:
		goto _out
	case 8:
		switch data[p] {
		case 9:
			goto tr14;
		case 32:
			goto tr14;
		case 33:
			goto tr15;
		case 44:
			goto tr16;
		case 59:
			goto tr17;
		case 124:
			goto tr15;
		case 126:
			goto tr15;
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 39:
				if 42 <= data[p] && data[p] <= 46 {
					goto tr15;
				}
			case data[p] >= 35:
				goto tr15;
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 94 <= data[p] && data[p] <= 122 {
					goto tr15;
				}
			case data[p] >= 65:
				goto tr15;
			}
		default:
			goto tr15;
		}
		goto tr12;
	case 9:
		switch data[p] {
		case 9:
			goto tr18;
		case 32:
			goto tr18;
		case 44:
			goto tr11;
		case 59:
			goto tr1;
		}
		goto tr0;
	case 1:
		switch data[p] {
		case 9:
			goto tr1;
		case 32:
			goto tr1;
		case 113:
			goto tr2;
		}
		goto tr0;
	case 2:
		if data[p] == 61 {
			goto tr3;
		}
		goto tr0;
	case 3:
		switch data[p] {
		case 48:
			goto tr5;
		case 49:
			goto tr6;
		}
		goto tr4;
	case 10:
		switch data[p] {
		case 9:
			goto tr19;
		case 32:
			goto tr19;
		case 44:
			goto tr20;
		case 46:
			goto tr21;
		}
		goto tr12;
	case 11:
		switch data[p] {
		case 9:
			goto tr22;
		case 32:
			goto tr22;
		case 44:
			goto tr11;
		}
		goto tr12;
	case 4:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr7;
		}
		goto tr4;
	case 12:
		switch data[p] {
		case 9:
			goto tr19;
		case 32:
			goto tr19;
		case 44:
			goto tr20;
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr7;
		}
		goto tr12;
	case 13:
		switch data[p] {
		case 9:
			goto tr19;
		case 32:
			goto tr19;
		case 44:
			goto tr20;
		case 46:
			goto tr23;
		}
		goto tr12;
	case 5:
		if data[p] == 48 {
			goto tr8;
		}
		goto tr4;
	case 14:
		switch data[p] {
		case 9:
			goto tr19;
		case 32:
			goto tr19;
		case 44:
			goto tr20;
		case 48:
			goto tr8;
		}
		goto tr12;
	case 6:
		if data[p] == 44 {
			goto tr10;
		}
		goto tr9;
	case 15:
		goto tr12;
	}

	tr12: cs = 0; goto _again
	tr0: cs = 0; goto f0
	tr4: cs = 0; goto f1
	tr1: cs = 1; goto _again
	tr17: cs = 1; goto f6
	tr2: cs = 2; goto _again
	tr3: cs = 3; goto _again
	tr21: cs = 4; goto _again
	tr23: cs = 5; goto _again
	tr9: cs = 6; goto _again
	tr11: cs = 7; goto _again
	tr16: cs = 7; goto f5
	tr20: cs = 7; goto f7
	tr15: cs = 8; goto _again
	tr13: cs = 8; goto f4
	tr18: cs = 9; goto _again
	tr14: cs = 9; goto f5
	tr5: cs = 10; goto f2
	tr22: cs = 11; goto _again
	tr19: cs = 11; goto f7
	tr7: cs = 12; goto _again
	tr6: cs = 13; goto f2
	tr8: cs = 14; goto _again
	tr10: cs = 15; goto f3

f2:
//line encoding.rl:31
 pb = p 
	goto _again
f0:
//line encoding.rl:41

			if pushed {
				ee = ee[:len(ee)-1]
			}
			p--
 cs = 6
goto _again

		
	goto _again
f6:
//line encoding.rl:49
 e.token = text() 
	goto _again
f3:
//line encoding.rl:69
 cs = 7
goto _again
 
	goto _again
f4:
//line encoding.rl:33

			e = encodingSpec{weight: 1}
			pushed = false
		
//line encoding.rl:31
 pb = p 
	goto _again
f1:
//line encoding.rl:47
 p--
 cs = 6
goto _again
 
//line encoding.rl:41

			if pushed {
				ee = ee[:len(ee)-1]
			}
			p--
 cs = 6
goto _again

		
	goto _again
f5:
//line encoding.rl:49
 e.token = text() 
//line encoding.rl:37

			ee = append(ee, e)
			pushed = true
		
	goto _again
f7:
//line encoding.rl:50
 e.weight, _ = strconv.ParseFloat(text(), 64) 
//line encoding.rl:37

			ee = append(ee, e)
			pushed = true
		
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
		switch _encodings_eof_actions[cs] {
		case 1:
//line encoding.rl:41

			if pushed {
				ee = ee[:len(ee)-1]
			}
			p--
 cs = 6
goto _again

		
		case 2:
//line encoding.rl:47
 p--
 cs = 6
goto _again
 
//line encoding.rl:41

			if pushed {
				ee = ee[:len(ee)-1]
			}
			p--
 cs = 6
goto _again

		
		case 6:
//line encoding.rl:49
 e.token = text() 
//line encoding.rl:37

			ee = append(ee, e)
			pushed = true
		
		case 8:
//line encoding.rl:50
 e.weight, _ = strconv.ParseFloat(text(), 64) 
//line encoding.rl:37

			ee = append(ee, e)
			pushed = true
		
//line encoding.go:408
		}
	}

	_out: {}
	}

//line encoding.rl:73


    if cs < encodings_first_final {
		return ee, fmt.Errorf("cs=%d data=%q p=%d", cs, data, p)
    }
    return ee, nil
}
