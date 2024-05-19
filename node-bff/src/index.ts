import cors from 'cors'
import express from 'express';
import routes from './routes';

const app = express();
const port = 8000;

// CORSミドルウェアの設定
app.use(cors({
    origin: 'http://localhost:3000' // ReactアプリがホストされているURL
}));

app.use('/api', routes);

app.listen(port, () => {
    console.log(`BFF service listening at http://localhost:${port}`);
});

export default app; // テストのためにエクスポート
