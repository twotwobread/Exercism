import subprocess
from pathlib import Path

FILE_PATH = Path(__file__).parent / "commands.txt"


def main():
    cmds = read_commands()
    execute_commands(cmds)


def read_commands():
    with open(FILE_PATH, "r") as file:
        return file.read().splitlines()


def execute_commands(cmds):
    for cmd in cmds:
        print(cmd)
        try:
            subprocess.run(cmd, shell=True, check=True)
        except Exception:
            pass


if __name__ == "__main__":
    main()
