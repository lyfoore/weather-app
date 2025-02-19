import React, { useState, useEffect } from 'react';

function App() {
  const [time, setTime] = useState('');

  useEffect(() => {
    fetch('/api/time')
      .then(response => response.json())
      .then(data => setTime(data.time))
      .catch(error => console.error('Error:', error));
  }, []);

  return (
    <div>
      <h1>Current time from backend:</h1>
      <p>{time || "Loading..."}</p>
    </div>
  );
}

export default App;