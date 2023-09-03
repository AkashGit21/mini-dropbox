import React from 'react'
import { useParams } from 'react-router-dom';


const FileView = ()  => {
  const { itemId } = useParams();

  return (
    <React.Fragment>
      <h2>Detailed File View</h2>

      <article className='file-view'>
        <section className='left'>
          <img className='file-icon-big' src='./images/file-logo.png' alt='file'/>
          <p className='filename'>File name</p>
        </section>
        <section className='right'>
          <p className='description'>
            <span className='title'>Description:</span>
            <span className='data'>Random description</span>
          </p>

          <p className='size'>
            <span className='title'>Size:</span>
            <span className='data size'> bytes</span>
          </p>

          <p className='updated-at'>
            <span className='title'>Last Updated:</span>
            <span className='data updated-at'>Time</span>
          </p>

          <div className='download btn'>
            <a href='#' className='btn-link'>Download</a>
          </div>

          <p className='adjacent-btns'>
            <div className='btn update'>
              <a href='#' className='btn-link'>Update</a>
            </div>
            <div className='btn delete'>
              <a href='#' className='btn-link'>Delete</a>
            </div>
          </p>
        </section>
      </article>
    </React.Fragment>
  )
}

export default FileView