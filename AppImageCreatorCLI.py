# nuitka-project: --standalone
# nuitka-project: --onefile
# nuitka-project: --include-data-files=default.png=default.png
# nuitka-project: --include-data-files=appimagetool.AppImage=appimagetool.AppImage
# nuitka-project: --windows-icon-from-ico=appimage.ico
# nuitka-project: --linux-icon=appimage.png
# nuitka-project: --output-filename=AppImageCreatorCLI
# nuitka-project: --windows-console-mode=disable
# nuitka-project: --remove-output
# nuitka-project: --quiet
import os, sys, subprocess
args = list(map(lambda arg: arg.split(':'), sys.argv[1:]))

def returnError(text):
    sys.exit(f"\033[91m{text}\033[0m")

def help():
    print("""
Usage: CreateAppImageCLI [-h] [-n|-N|--name:<name>] [-i|-I|--icon:<icon>] [-m|-M|--main:<main binary>] <data files>
    
Positional arguments:
    data files: You can have none of these, or infinite. These are the files your binary depends on.
                
Optional arguments:
    -h, -H, --help: Show this help message
    -n, -N, --name: Name of AppImage
    -i, -I, --icon: Icon for AppImage
    -m, -M, --main: Main Binary of AppImage
""".strip())
    sys.exit()

def main(name=os.path.basename(os.getcwd()), icon=os.path.join(os.path.dirname(os.path.abspath(__file__)), "default.png"), binary='', datafiles=[], nothing=True):
    print(f"Creating AppDir: {name}...")
    appDir = rf"{os.getcwd()}/{name}.AppDir"
    os.makedirs(f"{appDir}/usr/bin", exist_ok=True)
    open(rf"{appDir}/{name}.png", 'wb').write(open(icon, 'rb').read())
    open(rf"{appDir}/usr/bin/{name}", "wb").write(open(binary, 'rb').read() if binary else b'')
    with open(rf"{appDir}/{name}.desktop", "w") as f:
        f.write(f"""
[Desktop Entry]
Name={name}
Icon={name}
Type=Application
Categories=Utility
""".strip())
    with open(rf"{appDir}/AppRun", "w") as f:
        f.write(f"""
#!/bin/sh
                
this_dir="$(readlink -f "$(dirname "$0")")"
exec "$this_dir"/usr/bin/{name} "$@"
""".strip())
    for file in datafiles:
        open(f"{appDir}/{os.path.basename(file)}", 'wb').write(open(file, 'rb').read())
    print("Building AppImage...")
    result = None if nothing else subprocess.run(f"{os.path.join(os.path.dirname(os.path.abspath(__file__)), "appimagetool.AppImage")} {appDir} {name}.AppImage", shell=True, capture_output=True)
    subprocess.run(f"rm -rf {appDir}", shell=True)
    if result and result.returncode: returnError("ERROR: BUILD ERROR, CHECK IF FILE IN USE!")

if __name__=='__main__':
    name = os.path.basename(os.getcwd())
    icon = os.path.join(os.path.dirname(os.path.abspath(__file__)), "default.png")
    binary = ''
    datafiles = []
    if not args: main()
    for idx, arg in enumerate(arginitial[0] for arginitial in args):
        if arg in ['-h', '-H', '--help']: help()
        if arg in ['-n', '-N', '--name']: name = args[idx][1]
        if arg in ['-m', '-M', '--main']: binary = args[idx][1]
        if arg in ['-i', '-I', '--icon']: icon = args[idx][1] if icon.endswith(".png") else returnError("ERROR: FILE IS NOT '.png' IMAGE FILE!")
        if len(args[idx])==1: datafiles.append(arg)
    main(name, icon, binary, datafiles, nothing=False)