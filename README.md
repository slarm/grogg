# grogg
A small command line tool for testing grok patterns on file contents.

It uses the [vjeantet/grok](https://github.com/vjeantet/grok) library for parsing and [pterm](https://github.com/pterm/pterm) for some extra gloss.
# Usage
Available parameters are:
```
  -custom-file string
        File containing custom Grok patterns
  -failures
        Show only failures
  -input string
        Input file
  -line-length int
        Set max line length buffer in bytes (default 4096)
  -multiline string
        Multiline pattern
  -pattern string
        Grok pattern to match
  -prompt
        Prompt after each line
  -silent
        Show only summary at the end of execution
  -verbose
        Verbose mode
  -version
        Print version

```
# Examples
Show only lines where that didn't return any matches:
```
grogg -input=access.log -pattern='%{COMMONAPACHELOG}' -failures -prompt
```
Add additional patterns from file and match multiline events where every new events starts with a date on the format "mmm DD HH:MM:SS":
```
grogg -input=file.log -pattern='%{CUSTOMLOGFILE}' -custom-file=grok-patterns.txt -multiline='^\w{3}\s(\d{2}:){2}\d{2}'
```
