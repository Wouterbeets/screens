# screens
## Cycle screen configurations.

### requirements.
screens requires a .screens directory in your home where it can save the current state for a command.
In the directory, the filenames correspond to a screens arguments.
example of how to configure

```bash
mkdir ~/.screens
echo '{
	"Entries": [
		"xrandr --output eDP-1 --primary --mode 1920x1080 --pos 0x0 --rotate normal --output HDMI-3 --off --output HDMI-2 --off --output HDMI-1 --off --output DP-3 --off --output DP-2 --off --output DP-1 --off",
		"xrandr --output eDP-1 --off --output HDMI-3 --off --output HDMI-2 --off --output HDMI-1 --mode 1920x1200 --pos 0x0 --rotate normal --output DP-3 --off --output DP-2 --off --output DP-1 --off"
	],
	"Current": 0
}
' > ~/.screens/1screen

# now call the screens with the filename as an argument
screens 1screen
# screens inrcemented the index in the configuration file and the next time it's called executes the second entry
screens 1screen
```
### Complex commands
if you run complex commands its probably a good idea to run it with a script, my configuration file looks like this
```json
{
	"Entries": [
		"/home/bugless/.screens/scripts/1screen.sh",
		"/home/bugless/.screens/scripts/1screen_hdmi1.sh",
		"/home/bugless/.screens/scripts/1screen_hdmi2.sh"
	],
	"Current": 1
}
```


