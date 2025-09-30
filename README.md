# AppImageCreator
Creates AppImages using [AppImageTool](https://github.com/AppImage/appimagetool). Includes **AppImageCreatorCLI** and **AppImageCreatorGUI**

**Note: Please open issue if found**

## AppImageCreatorCLI
**This is a CLI version of AppImageCreator.**

Usage: `AppImageCreatorCLI.AppImage [-h] [-n|-N|--name:<name>] [-i|-I|--icon:<icon>] [-m|-M|--main:<main binary>] [-t|-T|--terminal] [-v|-V|--verbose] <data files>`

Creates the AppDir only if no arguments passed!

Positional arguments:<br>
    data files: You can have none of these, or infinite. These are the files your binary depends on.
                
Optional arguments:<br>
    -h, -H, --help: Show this help message<br>
    -n, -N, --name: Name of AppImage<br>
    -i, -I, --icon: Icon for AppImage<br>
    -m, -M, --main: Main Binary of AppImage<br>
    -t, -T, --terminal: Terminal app<br>
    -v, -V, --verbose: Verbose output

## AppImageCreatorGUI
**This is a GUI version of AppImageCreator**

Usage: Launch it and it'll work *OOTB*.