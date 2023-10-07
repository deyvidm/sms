remote := "root@${LINODE_IP}"

t:
	echo $(remote)
deploy:
	rsync -av -e ssh --exclude-from='.gitignore' . $(remote):/root/sms-frontend
