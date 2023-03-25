package wordmunger

var commonSubs = map[string][]string{
	"a": {"4", "@", "/-\\", "^", "∂", "α", "λ", "ά", "λ"},
	"b": {"6", "8", "|3", "13", "|o", "ß", "P>", "|:", "I3", "j3", "l3", "13", "I3", "β", "Ъ"},
	"c": {"(", "{", "<", "[", "©", "¢", "ç", "ć", "ċ"},
	"d": {"|)", "|>", "I>", "|}", "T)", "I7", "cl", "o|", "p>", "|o", "I7", "cl", "o|", "p>", "|o", "δ", "ď", "đ", "∂"},
	"e": {"3", "€", "ë", "&", "[-", "|=-", "έ", "ę", "ė", "ē", "є", "ѐ"},
	"f": {"|=", "ph", "|#", "|2", "ƒ"},
	"g": {"6", "9", "(_+", "ğ", "ĝ", "ǥ"},
	"h": {"|-|", "]-[", "(-)", ")-(", ":-:", "I-I", "I=I", "I^I", "I~I", "η", "ĥ", "ħ"},
	"i": {"1", "!", "ι", "ï", "ì", "ĭ", "í", "î", "ї"},
	"j": {"_|", "_/", "_7", "_)", "_(", "_>", "_<", "_|", "_/", "_7", "_)", "_(", "_>", "_<", "ј", "ĵ"},
	"k": {"|<", "|{", "|(", "|c", "|<", "|{", "|(", "|c", "|<", "|{", "|(", "|c", "|<", "|{", "|(", "|c", "κ", "ķ"},
	"l": {"1", "7", "|", "£", "1_", "|_", "ĺ", "ļ", "ł"},
	"m": {"|v|", "|\\/|", "м", "мм", "ɱ", "Μ"},
	"n": {"|\\|", "/\\/", "/V\\", "][\\][", "/\\", "л", "ń", "ň", "ņ"},
	"o": {"0", "()", "oh", "[]", "<>", "ό", "ö", "ó", "ø", "ð"},
	"p": {"|*", "|o", "|º", "|^(o)", "|>", "|\"", "|7", "ρ", "Þ"},
	"q": {"(_,)", "q", "φ", "Θ"},
	"r": {"|2", "|?", "/2", "lz", "2"},
	"s": {"5", "$", "z", "§", "ehs", "es"},
	"t": {"7", "+", "-|-", "']['", "†"},
	"u": {"|_|", "(_)", "L|", "µ"},
	"v": {"\\/", "\\/"},
	"w": {"\\/\\/", "\\^/", "\\|/", "\\V/", "\\X/", "\\|\\| /"},
	"x": {"%", "><", "}{"},
	"y": {"`/", "j", "`(", "7", "φ", "¥"},
	"z": {"2", "7_", ">_", "%", "s", "ehs", "es"},
}
