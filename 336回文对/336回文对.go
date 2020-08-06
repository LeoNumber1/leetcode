package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//workers := 3
	//ch := make(chan struct{})
	//worker := func() {
	//	// 干活干活干活
	//	rand.Seed(time.Now().UnixNano())
	//	i := rand.Intn(10)
	//	fmt.Printf("我开始工作了，%d秒\n", i)
	//	time.Sleep(time.Duration(i) * time.Second)
	//	ch <- struct{}{} // 通知管理者
	//}
	//leader := func() {
	//	cnt := 0
	//	for range ch {
	//		cnt++
	//		fmt.Println("====", cnt)
	//		if cnt == workers {
	//			break
	//		}
	//	}
	//	close(ch)
	//	// 检查工作成果
	//}
	//go leader()
	//for i := 0; i < workers; i++ {
	//	go worker()
	//}
	//
	//time.Sleep(10 * time.Second)
	//return

	words := []string{
		"abcd", "dcba", "lls", "s", "sssll",
		//"bat", "tab", "cat",
		//"", "",
		//"abc", "cba", "abc",
		//"abc",

		//"bbbbbabbabbbb", "baabbaa", "bbab", "bbbabbaaab", "abababbbbbab", "abb", "baaaabbb", "babbaaaba", "aab", "aaaab", "baabbbbabbaaaba", "baaab", "abbbab", "abaabbbabbabba", "aa", "aabbbaabba", "aaabbbbbaaabbbb", "bbaaaaba", "ababaaa", "aaaaa", "aaaaabbbbaaaba", "abbabbbaabbaabbb", "bbaba", "aaaaabbbabbbbaaaab", "abbbaa", "bbbabbaaa", "bbbaaabaabbbaaaaabaa", "aaaabbabb", "ababbababbbab", "aaaaababaababbbabaaa", "ba", "bbbbababbbabab", "baaaba", "aababbaaabbb", "aabbaaabbabaaababaab", "abbbb", "babaabaaababb", "bbbbabaaaab", "babbbbb", "babaaba", "aaba", "abababba", "a", "bb", "abaaab", "babbabaababbabaaba", "aaaaaababbbabaaabaa", "baabaaabb", "b", "bbaaaabbb", "abaaaaabaabbbaa", "ab", "bababaaaba", "aabababb", "ababaabbaababba", "bbb", "ababbaabababbbbbabb", "bbbbb", "abbbbaabaaaabb", "baba", "bbaabbabaaababaabbaa", "bbaabaabbabbaab", "bbbaabbab", "babbbbbaaaaabaa", "abbbbbbabbbabb", "abaa", "bbbbaababaab", "abaaababa", "aaaababaaababbaaba", "bbabbbabbbbbbaab", "abbabbabaabbabbbba", "abbbbbaabbbaaabaaaa", "bbaabababb", "aaabababaabbaaaaaaab", "ababbaabaaababb", "abbbbabaaabbaaabbab", "aababababbabaaa", "baabbaabbbaaaaaa", "bbbbbbbabbabbbbbb", "bbbabbabbabbabaabba", "babbbbabaaaabbabaab", "baabaabaabababaaabba", "bbaaaabbbbabbbaaaa", "aaaaabaabaa", "bbabaaabbbabaa", "baaabaaaaaab", "ababbbbbbbabaaaba", "abbbabaababbbbbaaa", "baaaaaabab", "aabbabba", "baaabbaabbbbaba", "aaaaabba", "babaaabbba", "bbbbab", "bbbbaabbaabab", "baa", "baababaa", "abbbbb", "babbaa", "abbbabbaa",

		//"hfjdhccfbfhhhghde", "jg", "dhjdjida", "f", "dediagbhecegdghaaaid", "icdabcebbhehhaahh", "eibedacciaa", "idfcdhgjgcjcgjig", "fhfidbggdghdibifeh", "bageigab", "eeahjicgdfgcbd", "gfeg", "g", "hbjgiabh", "cbhidai", "bachcbdcjegcbdijgdfd", "icjccejafcaeci", "cg", "geddabbahgcghgihcica", "bggadhcahhae", "diejjefbjeb", "eieahc", "igdbccch", "jfdhe", "j", "hchdaggbjhiccgach", "dhiad", "ehjdjciehifgbg", "hafggdfaeca", "b", "iddhacfijbiajffehi", "icihfiea", "cjdechbidfhged", "bbhdaecgabfega", "bdifcbbhdaaabjj", "dajfjhicfbddiciaeibi", "hfgchhbhbhdhajhbi", "gdifa", "bjdcefdjifjb", "abeajjadjdfii", "eeifadhabcfedbg", "gaeadgicfg", "hbjhcdaabaifahehfj", "iiaebcihfjeaaiijc", "hhgidag", "hecj", "hcgibfbihgacicaj", "aaabicdjb", "jjjchiacjicfiegggc", "bdaagfbeacjjh", "eebgbajgie", "aaedcicjfbggdjefaaff", "ggijjhecha", "dcd", "gceddjgjhfcfj", "giefff", "gibeceicg", "edfdcbjgchicghcf", "djjcjdajjhijjjdhh", "iggefcgbfg", "jfhjjdcfbefiggcdgi", "gadc", "iebagfeejchcd", "bhadjecgbfdfbabibadb", "fdeejfbhijdacfa", "gigdedjjbdedgifgghj", "ihhhhgib", "bbeiiadjbgjfajdbijah", "ac", "jiacaaebejabfdf", "fcjiihhhbjbabbdg", "fidgcgidaehbe", "ahehh", "ggigjddhajhdbafbahj", "fhicijfejaabdejea", "efjbgafebige", "dhdgfdcdfeifdjeebejd", "dfjcjjidcghii", "ddffjgeddg", "egjheejjfjghgha", "babgccegagcfidaaidd", "jacdab", "gfdbb", "fdhibbafeffgb", "i", "bfbbbhggcciie", "aajdicfbgbife", "deaeaafb", "effajbch", "ehiffebffadag", "ie", "ihdjjdieefgdbe", "igcidcdbae", "fee", "igfajj", "agebeadadbacbhghf", "abgeihfiihefdjjec", "cebfgbijbdie", "fcih", "jfagagaffbjhdhd", "aeaghaiidde", "aehegcc", "jjebb", "gjchcdbchcbceddgc", "bah", "jcfjfdjcjhijfdifad", "afdjiefjfcchcejhf", "ecjhijdheadfgjbadj", "c", "chfabceefhif", "fjfgcfhffbicjedcd", "hjidfic", "bhidheggaccjeaf", "baiciahegiabj", "haaffbhhacfha", "icdajb", "ghadbbdbigfdjcd", "gjdigffdhdjahfeja", "jgdbeia", "fcjdibbgaaceijcadje", "agaij", "dghahcfhabeccdjdii", "cfbebcggadbbaeidac", "dfbdejhdehahjjc", "aieabiafibidi", "ehe", "a", "jehidebjj", "acecjibeceec", "fdbhbhcf", "dfcdcicdeeijcghi", "ehghggeejccjhj", "agfgheea", "aiifigdbggef", "ffibh", "gbfc", "digaeag", "ijaeee", "icjacfd", "ccffcghdehgifcfijhd", "fhgijaia", "iegdhieghgjhajiibe", "ahfjcdibefdbgff", "ddgbgbdej", "fcifidi", "bdjddhcaieffega", "diegae", "ahgfbeegjhighfdh", "fbcjcdfifhh", "hcajij", "bcfbcbiafcjc", "egjhefbdibdcibjgah", "hdibeidhgfei", "ifjgcbhdcbceedi", "igjjcgd", "ijbgibfbfjhhhcgf", "hfdidbiee", "egaccifbgeeccbhce", "aggih", "aaabjehebgh", "deifdbcgi", "jcbcebdgehbgacfd", "dgadijgchfhfabaj", "acaiicgdeijajjbbjbih", "bfefheifed", "jjajccagjdebhbfi", "abeadgcgjhjbiahcjhij", "cigc", "d", "ibbdadbfiabacdaii", "gj", "edafihdjhagjicde", "ibgdahbdjebdhbjd", "deciacgbagfge", "bjdhcefhhj", "cdffb", "ghhfgadgeh", "dfbj", "dieefgjedbdbffgdih", "ehjda", "ffacja", "gbcbfgdb", "faiahggicabgfah", "ahg", "egibjhb", "fiiafeahaacbafie", "ghdgdccjagiajg", "edihdfgbadf", "bjh", "cfgebdga", "fbfaefbee", "gehchffaibjdgej", "iedf", "cggccdjf", "eejfabbaifjagigjhe", "biiedhcgcacef", "dfjj", "eeegbjgciabaiagafhf", "fhbac", "cebaeaed", "ii", "gcg", "hicfdbcjheabhe", "ejcceabdbaecajacia", "bebiaaab", "aibfjchihfca", "fhecfiadbdhhjgjd", "aieidehjjbcfhc", "fhhifibhfgbafffgghfj", "dfjfiejfhcifaia", "ddgeagibgjebehbe", "dhe", "fegdccjea", "afjcdbjgiedaahadada", "ebdefficffbabfjib", "hedcdajgjcjj", "adbfbbaag", "fehjhgg", "dfdidfh", "gbafihehchiiecfjd", "ehbbfbgbhfcij", "gaiaihbaebdbgagjhc", "bjgdcba", "dbdiidfacecgfia", "habbahi", "cicjch", "geia", "gedgfeaca", "ghjcdaicddecjejcjbbf", "ejecah", "cehiidhefebbcifhidh", "jfajahghhihgejif", "df", "ehc", "idecbgbbfcfagi", "agghc", "iebchg", "aaidgiebhgaegagejc", "ighbbgeebgedgg", "jfhdecjegbcaffedeeh", "ecdgfhdhcj", "geieedfhddbjacejifc", "afiachijidbgb", "jgdhebibagiab", "eibgdajdfafjhjbefjch", "gehhchbghcaefgc", "hhicchcaai", "bjfcehajfabg", "jgebijcebadfhedbjed", "hhjdgibhajgdj", "fjjbijcfaddc", "hab", "hiibbbcbafdgjdjjjfad", "cfafbgeeb", "ffai", "agiafgdhgbciaahcc", "bgidifibhjbchjec", "jjgegfchhjdhj", "afgc", "bjjdhdbai", "bgg", "gbhfbjf", "aeig", "hjfedhhbaejc", "gee", "ajigidibe", "iegchabgafdbgcfiebi", "ab", "jd", "jhifbhagdfe", "hbihfaeiibbiaj", "ecje", "fjbebd", "dhbahifecjagg", "hj", "aabjhcidbfdehigfdij", "hhghj", "icdacceh", "ddiejedbiiejdagcaac", "dbadbcjgjjjbchh", "jb", "fhdcijccffcaefeg", "hhaiihefecdjaicbj", "fjicca", "aagfaj", "deaaaf", "biejeg", "befjgigbbce", "bbccii", "igjbfihdhhcbjjdgddhb", "iaebdfghfdeh", "dihibgaijhbbdjjdfge", "egahhiiighaijgbidaba", "fef", "cfeifadjhdbfh", "bfdbjjjejcebjdbh", "dicggihe", "cicdjgeb", "fjgejhdchgjdgijbcai", "dcbeifdba", "hbggjacfjeffjhdhhj", "djjbicbijjea", "eahbejbjhaebc", "bhhgedehh", "dg", "jaagfcccjhjfbfdjd", "ffccceaheddjdfg", "bbacha", "hbcgfjfabifbhfbjed", "eejfficdd", "idid", "dhf", "ghffigiejjgdbjh", "efcebfhefiic", "fhhgbhiedghibecgj", "ca", "caaedegefdf", "ihcbjhhejfag", "iagifii", "bagbejjdg", "gffgihchfcecgebi", "dajfia", "fbaj", "bgefadaidfhcgcg", "jafdcgjaifibjc", "jbeaaigddj", "egjdhjaheifhecafaj", "ecfcfig", "dijdhf", "ijddacafafhfihi", "jhbfdhhbfd", "ehjaadchajceagg", "eagi", "ddhhicbg", "addhbd", "eehbibfigegffebhjbge", "beaegjeaibd", "aceaccbbgfecadihjee", "gccajbfjdh", "jeghcgjejhi", "gfdbijjbcefh", "iagaaijdiead", "gdcbbfcb", "dgihjaigadfbggjigeb", "dbcibjcfheffb", "gfgc", "iee", "ggecaadafidfafhfddig", "ahjeafafiadb", "jjijebjbcigiagacj", "idijaaefdfdheagabjjj", "cceegeajeagd", "dfagibjfddjga", "hbifcjceeecgfdgchfc", "jahfgg", "iejadfg", "eidegibhffdgh", "fcbaefhe", "ahffhgcagfjjib", "bcgfjigcbhehhdej", "e", "cbajcaijaabb", "jgfaiib", "jhcdejaeafceefbgej", "icbggbfdacjjchiegb", "bfehcb", "degajfegedj", "hjdcehdg", "gjf", "idafjaffihhf", "bffafeffedehagf", "icj", "ighjdibbdjidifgiibge", "fhhgghgcjcacje", "dcjgbgfbbjgjdfaij", "jajbbgcf", "efjhiejeedefcd", "biigiafgfccfgjbhc", "baieeiabjigiedih", "ejhd", "egadigfeaghcggdighf", "jhhei", "cbcbfgcffjae", "igjhbhgeeahjhgdfbe", "dcaigidfegdijgcbjc", "febjicbgai", "bgbgehigcfdei", "ghgh", "cahg", "fjbgbeife", "dgiaefhgbcgffb", "geigcbijaj", "ajggdbfa", "iebjfcjbfgiigac", "jbhicefbccieggccag", "aegababcgdecfjcgdghj", "eedafidgcghgggeec", "gjefjgihdcedfbjdce", "addeci", "ahagje",
	}
	//t1 := time.Now()
	//fmt.Println(len(palindromePairs(words)), "t1 =", time.Since(t1))
	//t2 := time.Now()
	//fmt.Println(len(palindromePairsSync(words)), "t2 =", time.Since(t2))
	t3 := time.Now()
	fmt.Println(len(palindromePairsOfficial(words)), "t3 =", time.Since(t3))
}

func palindromePairs(words []string) [][]int {
	res := make([][]int, 0)
	for k, v := range words {
		for j := k + 1; j < len(words); j++ {
			if isPalindromeString(v + words[j]) {
				res1 := make([]int, 0)
				res1 = append(res1, k)
				res1 = append(res1, j)
				res = append(res, res1)
			}
			if isPalindromeString(words[j] + v) {
				res1 := make([]int, 0)
				res1 = append(res1, j)
				res1 = append(res1, k)
				res = append(res, res1)
			}
		}
	}
	return res
}

func isPalindromeString(x string) bool {
	if len(x) == 0 {
		return true
	}
	n := len(x)
	for i := 0; i < n/2; i++ {
		if x[i] != x[n-1-i] {
			return false
		}
	}
	return true
}

func palindromePairsSync(words []string) [][]int {
	res := make([][]int, 0)
	n := len(words)
	if n < 2 {
		return res
	}

	wg := new(sync.WaitGroup)
	wg.Add(n * (n - 1))
	m := new(sync.RWMutex)

	for k, v := range words {
		for j := k + 1; j < len(words); j++ {
			go func(i1, i2 int, v1, v2 string) {
				defer wg.Done()
				if isPalindromeString(v1 + v2) {
					res1 := make([]int, 0)
					res1 = append(res1, i1)
					res1 = append(res1, i2)
					m.Lock()
					res = append(res, res1)
					m.Unlock()
				}
			}(k, j, v, words[j])

			go func(i1, i2 int, v1, v2 string) {
				defer wg.Done()
				if isPalindromeString(v1 + v2) {
					res1 := make([]int, 0)
					res1 = append(res1, i2)
					res1 = append(res1, i1)
					m.Lock()
					res = append(res, res1)
					m.Unlock()
				}
			}(k, j, words[j], v)
		}
	}
	wg.Wait()
	return res
}

func palindromePairsOfficial(words []string) [][]int {
	wordsRev := []string{}
	indices := map[string]int{}

	n := len(words)
	for _, word := range words {
		wordsRev = append(wordsRev, reverse(word))
	}
	for i := 0; i < n; i++ {
		indices[wordsRev[i]] = i
	}

	ret := [][]int{}
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		if m == 0 {
			continue
		}
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findWord(word, 0, j-1, indices)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, m-1, indices)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}

func findWord(s string, left, right int, indices map[string]int) int {
	if v, ok := indices[s[left:right+1]]; ok {
		return v
	}
	return -1
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}
