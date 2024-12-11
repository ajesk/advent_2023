DAY="$1"
SOURCE="day${DAY}.go"
TEST="day${DAY}_test.go"

cp -v base.go.template $SOURCE
cp -v base_test.go.template $TEST
sed -i "s/\[x\]/${DAY}/g" $SOURCE
sed -i "s/\[x\]/${DAY}/g" $TEST
touch ./data/day${DAY}.txt
touch ./data/day${DAY}demo.txt