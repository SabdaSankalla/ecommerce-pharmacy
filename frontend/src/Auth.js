const TOKEN_KEY = 'token';

export const login = (username, password) => {
  if (username === 'user' && password === 'password') {
    localStorage.setItem(TOKEN_KEY, 'your_token'); 
    return true;
  }
  return false;
};

export const logout = () => {
  localStorage.removeItem(TOKEN_KEY);
};

export const isAuthenticated = () => {
  return !!localStorage.getItem(TOKEN_KEY);
};