import React from 'react';
import LogsRow from './LogsRow'

export default function LogsTable(props) {
    return <table>
        <tr>
            <th>IP Address</th>
            <th>Accessed On</th>
        </tr>
        { props.urlLogs.map(
            (urlLog, i) => <LogsRow key={i} logInfo={urlLog}/>)}
    </table>
    
}