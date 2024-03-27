import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './VideoList.css'; // Import CSS file for styling

const VideoList = () => {
  const [videos, setVideos] = useState([]);

  useEffect(() => {
    const fetchVideos = async () => {
      try {
        const response = await axios.get('http://localhost:8000/videos');
        const videoData = response.data.map(item => ({
          id: item.id.videoId,
          title: item.snippet.title,
          description: item.snippet.description,
          thumbnail: item.snippet.thumbnails.medium.url
        }));
        setVideos(videoData);
      } catch (error) {
        console.error(error);
      }
    };

    fetchVideos();
  }, []);

  return (
    <div className="video-list-container">
      <h1 style={{ color: '#007bff' }}>Video Dashboard</h1>
      <div className="video-grid">
        {videos.map((video) => (
          <div key={video.id} className="video-item">
            <h2 style={{ color: '#333' }}>{video.title}</h2>
            <img src={video.thumbnail} alt={video.title} />
            <p style={{ color: '#666' }}>{video.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default VideoList;
