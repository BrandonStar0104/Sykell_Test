import React, { useState } from 'react';
import axios from 'axios';

function URLForm() {
  const [url, setUrl] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:8080/process', { url }, {
        headers: {
          'Authorization': `Bearer your-secure-token`,
          'Content-Type': 'application/json'
        }
      });
      setUrl('');
      // Redirect to URL list or show success message
    } catch (error) {
      console.error("Error processing URL:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="url"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        required
        placeholder="Enter URL"
      />
      <button type="submit">Process URL</button>
    </form>
  );
}

export default URLForm;
