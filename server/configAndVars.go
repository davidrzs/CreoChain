package server

// WelcomeCreo string
const WelcomeCreo = "   _____                  ____  _            _        _           _          _____                          \r\n  / ____|                |  _ \\| |          | |      | |         (_)        / ____|                         \r\n | |     _ __ ___  ___   | |_) | | ___   ___| | _____| |__   __ _ _ _ __   | (___   ___ _ ____   _____ _ __ \r\n | |    | '__/ _ \\/ _ \\  |  _ <| |/ _ \\ / __| |/ / __| '_ \\ / _` | | '_ \\   \\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|\r\n | |____| | |  __/ (_) | | |_) | | (_) | (__|   < (__| | | | (_| | | | | |  ____) |  __/ |   \\ V /  __/ |   \r\n  \\_____|_|  \\___|\\___/  |____/|_|\\___/ \\___|_|\\_\\___|_| |_|\\__,_|_|_| |_| |_____/ \\___|_|    \\_/ \\___|_|   \r\n                                                                                                            \r\n                           "

// SingleHashCheck is the result of a single has check.
type SingleHashCheck struct {
	hash1 string
	hash2 string
	same  bool
}

//HashResult is the struct we return as json when we check all hashes.
type HashResult struct {
	HashesOk      bool
	DiscrepancyID int
}

// BlockAdder is a help structure facilitating adding a block to a chain
type BlockAdder struct {
	Content        string
	Authentication string
}

/*
# General Comments:

We have a versioned API -> we are currently v1
*/
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

//AddBlock help us to püarse the JSON when we are adding a new block to the end of a chain.
type AddBlock struct {
	Data     string
	Authcode string
}
