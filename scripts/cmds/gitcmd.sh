git --version
git clone https://github.com/bgoldovsky/micro

git add .
git commit -m "fixed"
git push -u origin master

git stash
git pull
git stash pop