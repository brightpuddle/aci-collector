import sys

pwd = sys.argv[1]
new_pwd = ""

for a, b in zip(pwd, "_" * len(pwd)):
    new_pwd += a + b

print(new_pwd)
