import type { NextPage } from 'next'
import Script from 'next/script'
import Head from 'next/head'
import styles from '../styles/Home.module.css'
import * as React from 'react'
// @ts-ignore
import SwipeableViews from 'react-swipeable-views'
import Divider from '@mui/material/Divider'
import { useTheme } from '@mui/material/styles'
import AppBar from '@mui/material/AppBar'
import Grid from '@mui/material/Grid'
import Stack from '@mui/material/Stack'
import Tabs from '@mui/material/Tabs'
import Button from '@mui/material/Button'
import Tab from '@mui/material/Tab'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Link from 'next/link'

interface TabPanelProps {
  children?: React.ReactNode;
  dir?: string;
  index: number;
  value: number;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`full-width-tabpanel-${index}`}
      aria-labelledby={`full-width-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `full-width-tab-${index}`,
    'aria-controls': `full-width-tabpanel-${index}`,
  };
}

const Home: NextPage = () => {
  const theme = useTheme();
  const [value, setValue] = React.useState(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  const handleChangeIndex = (index: number) => {
    setValue(index);
  };

  return (
    <div className={styles.container}>
      <Script
        src="https://www.googletagmanager.com/gtag/js?id=G-DBZDFPJCJ9"
        strategy="afterInteractive"
      />
      <Script id="google-analytics" strategy="afterInteractive">
        {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){window.dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', 'G-DBZDFPJCJ9');
        `}
      </Script>
      
      <Head>
        <title>Shortlink | Landing</title>
        <meta name="description" content="Routing by project" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Grid
        container
        direction="row"
        justifyContent="center"
        alignItems="center"
      >
        <Box sx={{ bgcolor: 'background.paper', width: 700 }}>
          <AppBar position="static">
            <Tabs
              value={value}
              onChange={handleChange}
              indicatorColor="secondary"
              textColor="inherit"
              variant="fullWidth"
              aria-label="full width tabs example"
            >
              <Tab label="Application" {...a11yProps(0)} />
              <Tab label="Infrastructure" {...a11yProps(1)} />
              <Tab label="Docs" {...a11yProps(2)} />
            </Tabs>
          </AppBar>
          <SwipeableViews
            axis={theme.direction === 'rtl' ? 'x-reverse' : 'x'}
            index={value}
            onChangeIndex={handleChangeIndex}
          >
            <TabPanel value={value} index={0} dir={theme.direction}>
              <Typography variant="h5" component="h4" align={"center"}>
                Shortlink service (Microservice example)
              </Typography>

              <Stack
                spacing={{ xs: 1, sm: 2, md: 4 }}
                direction="row"
                divider={<Divider orientation="vertical" flexItem />}
                mt={2}
                justifyContent="center"
                alignItems="center"
              >
                <Link href="/next">
                  <Button variant="outlined">
                    Next
                  </Button>
                </Link>
              </Stack>
            </TabPanel>

            <TabPanel value={value} index={1} dir={theme.direction}>
              <Typography variant="h5" component="h4" align={"center"}>
                Infrastructure services
              </Typography>

              <Stack
                spacing={{ xs: 1, sm: 2, md: 4 }}
                direction="row"
                divider={<Divider orientation="vertical" flexItem />}
                mt={2}
                justifyContent="center"
                alignItems="center"
              >
                <Link href="/prometheus">
                  <Button variant="outlined">
                    Prometheus
                  </Button>
                </Link>

                <Link href="/grafana">
                  <Button variant="outlined">
                    Grafana
                  </Button>
                </Link>

                <Link href="/rabbitmq">
                  <Button variant="outlined">
                    RabbitMQ
                  </Button>
                </Link>

                <Link href="/kyverno/#/">
                  <Button variant="outlined">
                    Kyverno
                  </Button>
                </Link>
              </Stack>
            </TabPanel>

            <TabPanel value={value} index={2} dir={theme.direction}>
              <Typography variant="h5" component="h4" align={"center"}>
                Documentation and etc...
              </Typography>

              <Stack
                spacing={{ xs: 1, sm: 2, md: 4 }}
                direction="row"
                divider={<Divider orientation="vertical" flexItem />}
                mt={2}
                justifyContent="center"
                alignItems="center"
              >
                <Link href="https://github.com/batazor/shortlink">
                  <Button variant="outlined">
                    GitHub
                  </Button>
                </Link>

                <Link href="https://gitlab.com/shortlink-org/shortlink/">
                  <Button variant="outlined">
                    GitLab
                  </Button>
                </Link>

                <Link href="https://shortlink-org.gitlab.io/shortlink/">
                  <Button variant="outlined">
                    Swagger API
                  </Button>
                </Link>
              </Stack>
            </TabPanel>
          </SwipeableViews>
        </Box>
      </Grid>
    </div>
  )
}

export default Home
