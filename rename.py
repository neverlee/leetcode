
import os

root = "done/"
for fname in os.listdir(root):
    with open(root + fname) as f:
        fline = f.readline()
        fline = fline[3:].strip()
        if len(fname[:-3]) <= 4:
            print(root+fname, "-->>", root+fline+".rs")
            os.rename(root+fname, root+fline)
