import React, { useState } from 'react';
import { Box, Button, Input, Text, FormControl, FormLabel, FormErrorMessage } from '@chakra-ui/react';
import { login } from '../Auth'; // Impor fungsi login

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleLogin = () => {
    if (login(username, password)) {
      // Jika login berhasil, arahkan ke halaman utama
      window.location.href = '/'; // Ganti dengan navigasi yang sesuai
    } else {
      setError('Username atau password salah');
    }
  };

  return (
    <Box p={5}>
      <FormControl isInvalid={!!error}>
        <FormLabel>Username</FormLabel>
        <Input 
          value={username} 
          onChange={(e) => setUsername(e.target.value)} 
          placeholder="Masukkan username" 
        />
        <FormLabel>Password</FormLabel>
        <Input 
          type="password" 
          value={password} 
          onChange={(e) => setPassword(e.target.value)} 
          placeholder="Masukkan password" 
        />
        <FormErrorMessage>{error}</FormErrorMessage>
        <Button mt={4} onClick={handleLogin}>Login</Button>
      </FormControl>
    </Box>
  );
};

export default Login;