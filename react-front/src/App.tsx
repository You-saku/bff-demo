import { useEffect, useState } from 'react';

type Author = {
  id: number;
  name: string;
}

function App() {
  const [authors, setAuthors] = useState(Array <Author>());
  useEffect(() => {
    fetch('http://localhost:8000/api/authors')
        .then(response => response.json())
        .then(data => setAuthors(data))
    }, []);

return (
    <div>
        <h1>BFF Hello World</h1>
        <ul>
            {authors.map(author => (
                <li key={author.id}>{author.name}</li>
            ))}
        </ul>
    </div>
  );
}

export default App;
