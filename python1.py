# Open the file in read mode
file = open("example.txt", "r")

# Read all contents into memory
content = file.read()

# Print it
print(content)

# Always close when done
file.close()
