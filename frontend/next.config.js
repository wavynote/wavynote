/** @type {import('next').NextConfig} */

const nextConfig = {
    reactStrictMode: true,
    proxyOptions: {
        secure: false,
    },
    async rewrites() {
        return [
        {
            source: '/:path*',
            destination: `https://localhost:16770/wavynote/v1.0/:path*`,
        },
        ];
    },
};

export default nextConfig;
