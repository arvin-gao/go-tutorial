package gotutorial

const (
	_a = iota // 0
	_b = iota // 1
	_c = iota // 2
)

const (
	a = iota     // 0
	b            // 1
	c            // 2
	d = 5        // 5
	e            // 5
	f            // 5
	g = iota + 2 // 8
	h            // 9
)

const (
	Apple, Banana     = iota + 1, iota + 2 // 1, 2
	Cherimoya, Durian                      // 2, 3
	Elderberry, Fig                        // 3, 4

)

const (
	Open    = 1 << iota // 0001(1)
	Close               // 0010(2)
	Pending             // 0100(4)
)

const (
	_  = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota) // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
	PB                    // 1 << (10*5)
	EB                    // 1 << (10*6)
	ZB                    // 1 << (10*7)
	YB                    // 1 << (10*8)
)
