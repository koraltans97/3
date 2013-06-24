
// decode the regions+LUT pair into an uncompressed array
extern "C" __global__ void
regiondecode(float* __restrict__  dst, float* LUT, int8_t* regions, int N) {

    int i =  ( blockIdx.y*gridDim.x + blockIdx.x ) * blockDim.x + threadIdx.x;
    if (i < N) {

        dst[i] = LUT[regions[i]];

    }
}

