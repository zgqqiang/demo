const express = require('express');
const axios = require('axios');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.raw({ type: '*/*' }));

const serviceList = process.argv[2] ? process.argv[2].split(',') : ['servicea', 'serviceb'];
const service = process.argv[3] || 'servicea';
const code = parseInt(process.argv[4]) || 200;
const serverHandleTime = parseInt(process.argv[5]) || 0;
const printHeader = process.argv[6] === 'true';

app.use((req, res, next) => {
setTimeout(() => {
if (printHeader) {
console.log("请求头 =======");
console.log(req.headers);
console.log("请求头 =======");
}
next();
}, serverHandleTime * 1000);
});

app.all('/*', async (req, res) => {
const serviceIndex = serviceList.indexOf(service);

if (serviceIndex === -1 || serviceIndex === serviceList.length - 1) {
return res.status(code).send(process.env.POD_NAME);
}

const nextServiceName = serviceList[serviceIndex + 1];
const url = `http://${nextServiceName}:8080${req.originalUrl}`;

try {
const response = await axios({
method: req.method,
url,
data: req.body,
headers: req.headers,
});
res.status(code).send(`${process.env.POD_NAME}-->${response.data}`);
} catch (error) {
console.error(`请求失败，错误: ${error.message}`);
res.status(code).send(process.env.POD_NAME);
}
});

const PORT = 8080;
app.listen(PORT, () => {
console.log(`服务运行在 http://localhost:${PORT}`);
});