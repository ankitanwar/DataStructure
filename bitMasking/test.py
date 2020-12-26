
# Python 3 program to swap even  
# and odd bits of a given number 
  
# Function for swapping even  
# and odd bits 
def swapBits(x) : 
      
    # Get all even bits of x 
    even_bits = x & 0xAAAAAAAA
  
    # Get all odd bits of x 
    odd_bits = x & 0x55555555
      
    # Right shift even bits 
    even_bits >>= 1
      
    # Left shift odd bits 
    odd_bits <<= 1 
    print("Even and obits are",even_bits,odd_bits)
  
    # Combine even and odd bits 
    return (even_bits | odd_bits)  
  
  
# Driver program 
# 00010111 
x = 23
  
# Output is 43 (00101011) 
print(swapBits(x)) 
  