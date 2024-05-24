import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

function URLDetails() {
  const { id } = useParams();
  const [details, setDetails] = useState(null);

  useEffect(() => {
    const fetchDetails = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/result/${id}`, {
          headers: {
            'Authorization': `Bearer your-secure-token`,
            'Content-Type': 'application/json'
          }
        });
        setDetails(response.data);
      } catch (error) {
        console.error("Error fetching details:", error);
      }
    };

    fetchDetails();
  }, [id]);

  if (!details) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>Details for {details.URL}</h1>
      <p><strong>HTML Version:</strong> {details.html_version}</p>
      <p><strong>Page Title:</strong> {details.page_title}</p>
      <p><strong>Headings Count:</strong> {JSON.stringify(details.headings_count)}</p>
      <p><strong>Internal Links:</strong> {details.internal_links}</p>
      <p><strong>External Links:</strong> {details.external_links}</p>
      <p><strong>Inaccessible Links:</strong> {details.inaccessible_links}</p>
      <p><strong>Contains Login Form:</strong> {details.has_login_form ? 'Yes' : 'No'}</p>
    </div>
  );
}

export default URLDetails;
