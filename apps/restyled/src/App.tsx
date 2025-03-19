import styled from "styled-components";
import "./App.css";
import Button from "./ui/Button";

function App() {
    return (
        <Grid>
            <Button size="small" variant="fill">
                Button
            </Button>
            <Button size="small" variant="outline">
                Button
            </Button>
            <Button size="small" variant="ghost">
                Button
            </Button>
            <Button size="medium" variant="fill">
                Button
            </Button>
            <Button size="medium" variant="outline">
                Button
            </Button>
            <Button size="medium" variant="ghost">
                Button
            </Button>
            <Button size="large" variant="fill">
                Button
            </Button>
            <Button size="large" variant="outline">
                Button
            </Button>
            <Button size="large" variant="ghost">
                Button
            </Button>
        </Grid>
    );
}

const Grid = styled.div`
    display: grid;
    grid-template-columns: repeat(3, 180px);
    gap: 64px;
    padding: 48px;
    place-items: center;
`;

export default App;
