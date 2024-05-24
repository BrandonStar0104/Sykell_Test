import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

function URLList() {
  const [urls, setUrls] = useState([]);

  useEffect(() => {
    const fetchUrls = async () => {
      try {
        const response = await axios.get('http://localhost:8080/results', {
          headers: {
            'Authorization': `Bearer your-secure-token`,
            'Content-Type': 'application/json'
          }
        });
        setUrls(response.data);
      } catch (error) {
        console.error("Error fetching URLs:", error);
      }
    };

    fetchUrls();
  }, []);

  return (
    <div>
      <h1>Processed URLs</h1>
      <ul>
        {urls.map((url) => (
          <li key={url.url}>
            <Link to={`/details/${url.url}`}>{url.url}</Link> 
          </li>
        ))}
      </ul>
    </div>
  );
}

export default URLList;
