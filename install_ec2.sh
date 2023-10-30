####
# web frontend
####
cd web-front
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash
. ~/.nvm/nvm.sh
nvm install --lts
npm i


####
# dispatcher and web server run on Go
# they also need Redis
####
sudo yum install golang -y
cd dispatcher
make build 

cd ../web-server
make build

sudo dnf install -y redis6
sudo systemctl start redis6
sudo systemctl enable redis6
sudo systemctl is-enabled redis6
redis6-server --version
redis6-cli ping

####
# Extra stuff to get this whole blob workin
###
sudo yum install -y nginx
sudo systemctl enable nginx
sudo systemctl start nginx
sudo systemctl status nginx

