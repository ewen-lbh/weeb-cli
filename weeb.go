package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Convenience for weebs using the command line.

Usage:
	weeb SERVICE watch SHOW [EPISODE]
	weeb SERVICE watchlist [STATE] [--open]
	weeb SERVICE read [TITLE]
	weeb SERVICE news [SEARCH-TERM...]
	weeb SERVICE (--version | --about | --help)

Configuring:
	SERVICE allows you to choose which service to use. The supported values: MAL, AniList and Kitsu.
	To avoid typing "weeb AniList watch hyouka" each and every time, use an alias:
	
	alias weeb='weeb AniList'
	
	So that you only need to type 'weeb'
	
Options:
	-o --open             (command: watchlist) Open in your web browser instead of showing it on the terminal
	   --version          Print weeb-cli's version
	   --about            Print some useless info about the guy who made this stupid CLI
	   --help             Print this help
	
Commands:
	watch  
		Watch the next episode of SHOW, or watch episode EPISODE if specified, and update your watchlist. 
		
		Your watchlist is updated as follows:
			- Mark the show as "Watching"
			- Set the specified episode to "Watched"
		Of course, how these concepts are actually implemented vary depending on which service you use, 
		in most cases, you can't mark a specific episode as "watched" without marking all of the previous ones as well.
		
		EPISODE can use the following formats:
		[eE]?([\d\.]+)             - Watch episode of currently-watching season eg. weeb watch hyouka 5
		[sS]([\d\.]+)              - Watch first episode of season              eg. weeb watch kaguya-sama S2
		[sS]([\d\.]+)[eE]([\d\.]+) - Watch episode of season                    eg. weeb watch kaguya-sama S2E4
	
	read
		Read the next chapter of TITLE, or read CHAPTER if sepcified, and update your watchlist
		
		The watchlist is updated in the same way as with the 'watch' command.
		
		CHAPTER uses the following format:
		(?:ch|CH|Ch)?([\d\.]+)     - Read chapter, eg. weeb read 'domestic na kanojo' ch37
	
	watchlist
		This command shows your watchlist in your terminal.
		
		Use STATE to show only a section of your watchlist.
		Possible values:
		- pta, planning, want, planned, plan    -> "Plan to watch" section
		- current, currently, reading, watching -> "Watching" section
		- paused, pause, hold, onhold, on-hold  -> "On Hold" section
		- dropped, abandoned, given-up          -> "Dropped" section
	
	news
		Aggregate RSS feeds from multiple anime news sites and show them in your terminal.
		You can use SEARCH-TERM... to search for news about specific things, eg:
		
		weeb news yakusoku no neverland release date
`
	args, _ := docopt.ParseDoc(usage)
	fmt.Println(args)
}
