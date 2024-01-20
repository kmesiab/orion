import React from 'react';
import { parseDiff, Diff, Hunk } from 'react-diff-view';
import 'react-diff-view/style/index.css'; // Ensure you import the default styles which you can then override

const DiffViewer = ({ text }) => {
    let diffContent;

    if (text) {
        const files = parseDiff(text);
        diffContent = files.map((file, index) => (
            <div key={index} className="file-diff">
                <h3>{file.newPath}</h3>
                <Diff viewType="split" diffType={file.type} hunks={file.hunks}>
                    {hunks => hunks.map(hunk => (
                        <Hunk key={hunk.content} hunk={hunk}>
                            {(hunk, hunkIndex) => (
                                // Map over the changes within the hunk
                                hunk.changes.map((change, changeIndex) => (
                                    <div key={hunkIndex + '-' + changeIndex}
                                         className={`line ${change.type === 'insert' ? 'diff-code-insert' : change.type === 'delete' ? 'diff-code-delete' : ''} code`}>
                                        {change.content}
                                    </div>
                                ))
                            )}
                        </Hunk>
                    ))}
                </Diff>
            </div>
        ));
    } else {
        diffContent = <div>No diff to display</div>;
    }

    return (
        <div className="diff-viewer">
            {diffContent}
        </div>
    );
};

export default DiffViewer;
