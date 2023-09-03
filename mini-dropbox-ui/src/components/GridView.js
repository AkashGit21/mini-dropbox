import React, { useState,useEffect} from 'react';
import { useHistory } from 'react-router-dom';
import FileView from './FileView';
// import { apiHost } from './config';

const GridView = () =>{
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const history = useHistory();

  useEffect(() => {
    // const apiURL = apiHost + '/api/files';
    const apiURL = 'http://localhost:8082/api/files';
    console.log("API URL: " + apiURL)
    // Replace 'your-api-endpoint' with the actual API endpoint
    fetch(apiURL)
      .then((response) => response.json())
      .then((data) => {
        setItems(data);
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  function openFileView(id) {
    console.log("File ID is: "+ id);
    history.push(`/files/${id}`);
    // <Link to="/api/files/id"></Link>
  }

  return (
    <React.Fragment>
      <section className='menu'>
        <h2 className='heading'>User Files</h2>
        <div className='upload-btn'>Upload Files</div>
      </section>
      <div className="grid-container">
        {files.map((item) => (
          <div key={item.id} className="grid-item" onClick={() => openFileView(item.id)}>
            <img className='file-icon' src='./images/file-logo.png' alt='file'/>
            <span className='file-name'>{item.filename}</span>
          </div>
        ))}
      </div>
    </React.Fragment>
  );
}

export default GridView;
