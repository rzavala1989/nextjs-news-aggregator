// src/utils/auth.js
// import useUserStore from '../store/userStore';
//
// export const login = async (username, password, showToast) => {
//   try {
//     const response = await fetch('/login', {
//       method: 'POST',
//       headers: {
//         'Content-Type': 'application/json',
//       },
//       body: JSON.stringify({ username, password }),
//     });
//
//     if (response.ok) {
//       const data = await response.json();
//       useUserStore.getState().setUser(data.user);
//       showToast('Login successful', 'success');
//     } else {
//       showToast('Login failed', 'error');
//     }
//   } catch (error) {
//     showToast('An error occurred during login', 'error');
//   }
// };
//
// export const register = async (username, password, showToast) => {
//   try {
//     const response = await fetch('/register', {
//       method: 'POST',
//       headers: {
//         'Content-Type': 'application/json',
//       },
//       body: JSON.stringify({ username, password }),
//     });
//
//     if (response.ok) {
//       const data = await response.json();
//       useUserStore.getState().setUser(data.user);
//       showToast('Registration successful', 'success');
//     } else {
//       showToast('Registration failed', 'error');
//     }
//   } catch (error) {
//     showToast('An error occurred during registration', 'error');
//   }
// };
