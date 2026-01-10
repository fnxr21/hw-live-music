echo 'Enter the file name'

read 'file'

git add "$file"

# git add .

echo 'Enter the commit message:'

read commitMessage

git commit -m "$commitMessage"

git push origin dev

# echo 'Enter the branch name:'

# git branch -v

# read branch

# git remote -v

# echo 'Enter the remote name:'

# read remote

# git push origin "${branch}"
# git push "${remote}" "${branch}"

echo 'Thank you for updated your repository'