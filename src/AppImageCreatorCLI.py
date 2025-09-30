# nuitka-project: --quiet
# nuitka-project: --onefile
# nuitka-project: --standalone
# nuitka-project: --remove-output
# nuitka-project: --linux-icon=../icons/appimage.png
# nuitka-project: --windows-console-mode=disable
# nuitka-project: --windows-icon-from-ico=../icons/appimage.ico
# nuitka-project: --output-filename=AppImageCreatorCLI
# nuitka-project: --include-data-files=../assets/default.png=default.png
# nuitka-project: --include-data-files=../assets/appimagetool.AppImage=appimagetool.AppImage
import os, sys, shutil, subprocess
args = [arg.split(':') for arg in sys.argv[1:]]
returnError = lambda text: sys.exit(f"\033[91m{text}\033[0m")
returnWarning = lambda text: print(f"\033[93m{text}\033[0m")

def help():
    print('\n'.join([
        "Usage: CreateAppImageCLI [-h] [-n|-N|--name:<name>] [-i|-I|--icon:<icon>] [-m|-M|--main:<main binary>] <data files>",
        "",
        "Positional arguments:",
        "\tdata files: These are the files your binary depends on.",
        "",
        "Optional arguments:",
        "\t-h, -H, --help: Show this help message",
        "\t-n, -N, --name: Name of AppImage",
        "\t-i, -I, --icon: Icon for AppImage",
        "\t-m, -M, --main: Main Binary of AppImage",
        "\t-t, -T, --terminal: Whether or not the app runs in a terminal",
    ]))
    sys.exit()

def main(name, icon, binary, datafiles, terminal, verbose, nothing):
    print(f"Creating AppDir: {name}...")
    appDir = rf"{os.getcwd()}/{name}.AppDir"
    os.makedirs(f"{appDir}/usr/bin", exist_ok=True)
    open(rf"{appDir}/{name}.png", 'wb').write(open(icon, 'rb').read())
    open(rf"{appDir}/usr/bin/{name}", "wb").write(open(binary, 'rb').read() if binary else b'')
    with open(rf"{appDir}/{name}.desktop", "w") as f:
        f.write('\n'.join([
            "[Desktop Entry]",
            "Version=1.0",
            f"Name={name}",
            f"Icon={name}",
            f"Exec={name}",
            "Type=Application",
            f"Terminal={str(terminal).lower()}",
            "Categories=Utility",
        ]))
    with open(rf"{appDir}/AppRun", "w") as f:
        f.write('\n'.join([
            '#!/bin/sh',
            f'this_dir="$(readlink -f "$(dirname "$0")")"',
            f'exec "$this_dir"/usr/bin/{name} "$@"',
        ]))
    [shutil.copy2(f, os.path.join(appDir, os.path.basename(f))) for f in datafiles]
    if not nothing: 
        print("Building AppImage...")
        result = subprocess.run(f"{os.path.join(os.path.dirname(os.path.abspath(__file__)), "appimagetool.AppImage")} {appDir} {name}.AppImage", shell=True, capture_output=not verbose)
        shutil.rmtree(appDir)
        returnError("ERROR: BUILD ERROR, CHECK IF FILE IN USE!") if result.returncode else print("SUCCESS: AppImage Built Successfully!")

if __name__=='__main__':
    name = os.path.basename(os.getcwd())
    icon = os.path.join(os.path.dirname(os.path.abspath(__file__)), "default.png")
    binary = ''
    terminal = False
    verbose = False
    datafiles = []
    unknownFlags = []
    nothing = False if args else True
    for token in args:
        match token:
            case ["-h"] | ["-H"] | ["--help"]:
                help()
            case ["-t"] | ["-T"] | ["--terminal"]:
                terminal = True
            case ["-v"] | ["-V"] | ["--verbose"]:
                verbose = True
            case ["-n", value] | ["-N", value] | ["--name", value]:
                name = value
            case ["-m", value] | ["-M", value] | ["--main", value]:
                binary = value
            case ["-i", value] | ["-I", value] | ["--icon", value]:
                icon = value if value.endswith('.png') else returnError("ERROR: FILE IS NOT '.png' IMAGE FILE!")
            case [single]:
                datafiles.append(single)
    if unknownFlags: returnWarning(f"WARNING: Unknown Flags: {' '.join(unknownFlags)}")
    main(name, icon, binary, datafiles, terminal, verbose, nothing)