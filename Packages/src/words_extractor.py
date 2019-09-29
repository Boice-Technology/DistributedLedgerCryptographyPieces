f = open("word_database.csv","r")
lines = f.readlines()
f.close()
words = []
i = 0
for line in lines:
	if i<2048:
		k = line.split(" ")	
		words.append(k[1])
	else:
		break
	i+=1
f = open("MnemonicWords.go","w")
fileString = ""
for word in words:
	fileString += "\"" + word + "\", "
f.write(fileString)
f.close()


	