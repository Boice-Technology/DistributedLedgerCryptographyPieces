import pandas as pd

df = pd.read_csv("vocabwords.csv")

str = "\"" + df["Word"][0] + "\""
for i in range(1,2048):
	str = str + ", " + "\"" + df["Word"][i] + "\""

f = open("MnemonicWords.go","w")

f.write(str)
f.close()



