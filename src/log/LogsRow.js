import React from 'react';

export default function LogsTable(props) {
    return <tr>
        <td>{props.logInfo.ip_address}</td>
        <td>{props.logInfo.accessed_on}</td>
    </tr>
}