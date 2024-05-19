import { Router, Request, Response } from 'express';
import axios from 'axios';

const router = Router();

router.get('/authors', async (req: Request, res: Response) => {
    try {
        const response = await axios.get('http://localhost:80/api/authors');
        console.log('call go api');
        res.json(response.data);
    } catch (error) {
        res.status(500).json({ error: 'Failed to fetch data from backend' });
    }
});

export default router;
