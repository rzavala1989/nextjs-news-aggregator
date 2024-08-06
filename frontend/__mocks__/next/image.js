import React from 'react';

const Image = ({ src, alt, width, height, className }) => {
  return <img src={src} alt={alt} width={width} height={height} className={className} />;
};

export default Image;
