echo "Compiling warroom..."
rm -rf build/*
go build -C src -o ../build/warroom
echo "Warroom compiled!"
