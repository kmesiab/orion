import React, { useState } from 'react';
import { parseDiff, Diff, Hunk } from 'react-diff-view';
import { TabContext, TabList, TabPanel } from '@mui/lab';
import Tab from '@mui/material/Tab';
import Box from '@mui/material/Box';
import 'react-diff-view/style/index.css';

const FileDiffTabPanel = ({ file }) => (
    <div className="file-diff">
        <h3>{file.newPath}</h3>
        <Diff viewType="split" diffType={file.type} hunks={file.hunks}>
            {hunks => hunks.map((hunk, hunkIndex) => (
                <Hunk key={hunk.content} hunk={hunk}>
                    {hunk => hunk.changes.map((change, changeIndex) => (
                        <div key={`${hunkIndex}-${changeIndex}`}
                             className={`line ${change.type === 'insert' ? 'diff-code-insert' : change.type === 'delete' ? 'diff-code-delete' : ''} code`}>
                            {change.content}
                        </div>
                    ))}
                </Hunk>
            ))}
        </Diff>
    </div>
);

const DiffViewer = ({ text }) => {
    const [value, setValue] = useState('0');
    const files = parseDiff(text);

    if (!text) {
        return <div className="diff-viewer">Enter a GitHub pull request URL 👆</div>;
    }

    return (
        <Box sx={{ display: 'flex', height: '100%' }}>
            <TabContext value={value}>
                <Box sx={{ borderRight: 1, borderColor: 'divider', width: '20%', maxWidth: 300 }}>
                    <TabList
                        orientation="vertical"
                        onChange={(event, newValue) => setValue(newValue)}
                        aria-label="File diffs"
                    >
                        {files.map((file, index) => (
                            <Tab label={file.newPath} value={String(index)} key={index} />
                        ))}
                    </TabList>
                </Box>
                <Box sx={{ flexGrow: 1 }}>
                    {files.map((file, index) => (
                        <TabPanel value={String(index)} key={index}>
                            <FileDiffTabPanel file={file} />
                        </TabPanel>
                    ))}
                </Box>
            </TabContext>
        </Box>
    );
};

export default DiffViewer;
