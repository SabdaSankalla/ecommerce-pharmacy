import React from 'react';
import { Box, Heading, Button } from '@chakra-ui/react';
import { logout } from '../Auth';

const ItemList = () => {
  const handleLogout = () => {
    logout();
    window.location.href = '/login';
  };

  return (
    <Box>
      <Heading>Daftar Item</Heading>
      {/* Konten daftar item */}
      <Button onClick={handleLogout}>Logout</Button>
    </Box>
  );
};

export default ItemList;