// client/src/App.tsx
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Message } from '../../shared/src';

function App() {
  const [message, setMessage] = useState<string>('');

  useEffect(() => {
    axios.get<Message>('http://localhost:5000/api')
      .then(response => {
        setMessage(response.data.message);
      })
      .catch(error => {
        console.error("There was an error fetching the data", error);
      });
  }, []);

  return (
    {message}
  );
}

export default App;
