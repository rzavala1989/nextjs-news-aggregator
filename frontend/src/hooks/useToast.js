// src/hooks/useToast.js
import { toast } from 'react-toastify';

const useToast = () => {
  const showToast = (message, type = 'success') => {
    if (type === 'success') {
      toast.success(message);
    } else if (type === 'error') {
      toast.error(message);
    } else {
      toast.info(message);
    }
  };

  return {
    showToast,
  };
};

export default useToast;
