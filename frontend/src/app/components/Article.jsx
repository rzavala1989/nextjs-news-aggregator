import Image from 'next/image';
import PlaceholderImage from './PlaceholderImage';

const Article = ({ article }) => {
  const { title, content, author, url, image_url, published_at } = article;

  const truncateContent = (text, length) => {
    return text.length > length ? text.substring(0, length) + '...' : text;
  };

  const isValidImageUrl = (url) => {
    return url && url.startsWith('http');
  };

  return (
    <li className="card bordered shadow-lg flex flex-col md:flex-row items-start p-4 gap-4">
      <div className="flex-shrink-0 w-[130px] h-[130px] overflow-hidden rounded">
        {isValidImageUrl(image_url) ? (
          <Image
            src={image_url}
            alt={title}
            width={130}
            height={130}
            objectFit="cover"
            className="rounded"
          />
        ) : (
          <PlaceholderImage />
        )}
      </div>
      <div className="flex-grow">
        <h2 className="card-title text-xl font-semibold">{title}</h2>
        <p className="text-sm text-gray-500">{new Date(published_at).toLocaleDateString()}</p>
        <p className="mt-2">{truncateContent(content, 200)}</p>
        <a
          href={url}
          target="_blank"
          rel="noopener noreferrer"
          className="text-blue-500 hover:underline mt-2 block"
        >
          Read more
        </a>
      </div>
    </li>
  );
};

export default Article;
