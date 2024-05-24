import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import URLForm from './components/URLForm';
import URLList from './components/URLList';
import URLDetails from './components/URLDetails';
import './App.css';

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/" element={<URLForm />} />
          <Route path="/urls" element={<URLList />} />
          <Route path="/details/:id" element={<URLDetails />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
 