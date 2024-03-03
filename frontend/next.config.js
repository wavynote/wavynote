/** @type {import('next').NextConfig} */

const nextConfig = {
    reactStrictMode: false,
    // serverRuntimeConfig: {
    //   https: process.env.NEXT_PUBLIC_HTTPS === 'true',
    //   key: process.env.NEXT_PUBLIC_SSL_KEY,
    //   cert: process.env.NEXT_PUBLIC_SSL_CERT,
    // },
    async rewrites() {
        return [
        {
          source: '/:path*',
          destination: 'https://abc-wavynote.koyeb.app/:path*',
        },
      ];
    },
    async redirects() {
      return [
        {
          source: '/',
          destination: '/intro',
          permanent: false,
        },
      ]
    },
    // async headers() {
    //   return [
    //     {
    //       // matching all API routes
    //       source: "/:path*",
    //       headers: [
    //         { key: "Access-Control-Allow-Credentials", value: "true" },
    //         { key: "Access-Control-Allow-Origin", value: "https://localhost:16770/wavynote/v1.0/" },
    //         { key: "Access-Control-Allow-Methods", value: "GET,OPTIONS,PATCH,DELETE,POST,PUT" },
    //         { key: "Access-Control-Allow-Headers", value: "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version" },
    //       ],
    //     }
    //   ]
    // },
};

export default nextConfig;
