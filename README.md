# AppImageCreator
Creates AppImages using [AppImageTool](https://github.com/AppImage/appimagetool). Includes **AppImageCreatorCLI** and **AppImageCreatorGUI**

## AppImageCreatorCLI
**This is a CLI version of AppImageCreator.**

Usage: AppImageCreatorCLI.AppImage [-h] [-n|-N|--name:<name>] [-i|-I|--icon:<icon>] [-m|-M|--main:<main binary>] <data files>

Creates the AppDir only if no arguments passed!

Positional arguments:
    data files: You can have none of these, or infinite. These are the files your binary depends on.
                
Optional arguments:
    -h, -H, --help: Show this help message
    -n, -N, --name: Name of AppImage
    -i, -I, --icon: Icon for AppImage
    -m, -M, --main: Main Binary of AppImage

## AppImageCreatorGUI
**This is a GUI version of AppImageCreator**

Usage: Launch it and it'll work *OOTB*.